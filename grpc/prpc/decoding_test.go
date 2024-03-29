// Copyright 2016 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package prpc

import (
	"bytes"
	"context"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc/metadata"

	"go.chromium.org/luci/common/clock/testclock"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestDecoding(t *testing.T) {
	t.Parallel()

	Convey("readMessage", t, func() {
		var msg HelloRequest
		read := func(contentType string, body []byte) *protocolError {
			req := &http.Request{
				Body:   ioutil.NopCloser(bytes.NewBuffer(body)),
				Header: http.Header{},
			}
			req.Header.Set("Content-Type", contentType)
			return readMessage(req, &msg)
		}

		testLucy := func(contentType string, body []byte) {
			err := read(contentType, body)
			So(err, ShouldBeNil)
			So(msg.Name, ShouldEqual, "Lucy")
		}

		Convey("binary", func() {
			testMsg := &HelloRequest{Name: "Lucy"}
			body, err := proto.Marshal(testMsg)
			So(err, ShouldBeNil)

			Convey(ContentTypePRPC, func() {
				testLucy(ContentTypePRPC, body)
			})
			Convey(mtPRPCBinary, func() {
				testLucy(mtPRPCBinary, body)
			})
			Convey("malformed body", func() {
				err := read(mtPRPCBinary, []byte{0})
				So(err, ShouldNotBeNil)
				So(err.status, ShouldEqual, http.StatusBadRequest)
			})
			Convey("empty body", func() {
				err := read(mtPRPCBinary, nil)
				So(err, ShouldBeNil)
			})
		})

		Convey("json", func() {
			body := []byte(`{"name": "Lucy"}`)
			Convey(ContentTypeJSON, func() {
				testLucy(ContentTypeJSON, body)
			})
			Convey(mtPRPCJSONPBLegacy, func() {
				testLucy(mtPRPCJSONPB, body)
			})
			Convey("malformed body", func() {
				err := read(mtPRPCJSONPB, []byte{0})
				So(err, ShouldNotBeNil)
				So(err.status, ShouldEqual, http.StatusBadRequest)
			})
			Convey("empty body", func() {
				err := read(mtPRPCJSONPB, nil)
				So(err, ShouldNotBeNil)
				So(err.status, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("text", func() {
			Convey(mtPRPCText, func() {
				body := []byte(`name: "Lucy"`)
				testLucy(mtPRPCText, body)
			})
			Convey("malformed body", func() {
				err := read(mtPRPCText, []byte{0})
				So(err, ShouldNotBeNil)
				So(err.status, ShouldEqual, http.StatusBadRequest)
			})
			Convey("empty body", func() {
				err := read(mtPRPCText, nil)
				So(err, ShouldBeNil)
			})
		})

		Convey("unsupported media type", func() {
			err := read("blah", nil)
			So(err, ShouldNotBeNil)
			So(err.status, ShouldEqual, http.StatusUnsupportedMediaType)
		})
	})

	Convey("parseHeader", t, func() {
		c := context.Background()

		header := func(name, value string) http.Header {
			return http.Header{
				name: []string{value},
			}
		}

		Convey(HeaderTimeout, func() {
			Convey("Works", func() {
				now := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
				c, _ = testclock.UseTime(c, now)

				var err error
				c, err = parseHeader(c, header(HeaderTimeout, "1M"))
				So(err, ShouldBeNil)

				deadline, ok := c.Deadline()
				So(ok, ShouldBeTrue)
				So(deadline, ShouldHappenWithin, time.Second, now.Add(time.Minute))
			})

			Convey("Fails", func() {
				c, err := parseHeader(c, header(HeaderTimeout, "blah"))
				So(c, ShouldEqual, c)
				So(err, ShouldErrLike, HeaderTimeout+` header: unit is not recognized: "blah"`)
			})
		})

		Convey("Content-Type", func() {
			c, err := parseHeader(c, header("Content-Type", "blah"))
			So(err, ShouldBeNil)
			_, ok := metadata.FromIncomingContext(c)
			So(ok, ShouldBeFalse)
		})

		Convey("Accept", func() {
			c, err := parseHeader(c, header("Accept", "blah"))
			So(err, ShouldBeNil)
			_, ok := metadata.FromIncomingContext(c)
			So(ok, ShouldBeFalse)
		})

		Convey("Unrecognized headers", func() {
			test := func(c context.Context, header http.Header, expectedMetadata metadata.MD) {
				c, err := parseHeader(c, header)
				So(err, ShouldBeNil)
				md, ok := metadata.FromIncomingContext(c)
				So(ok, ShouldBeTrue)
				So(md, ShouldResemble, expectedMetadata)
			}

			headers := http.Header{
				"X": []string{"1"},
				"Y": []string{"1", "2"},
			}

			Convey("without metadata in context", func() {
				test(c, headers, metadata.MD(headers))
			})

			Convey("with metadata in context", func() {
				c = metadata.NewIncomingContext(c, metadata.MD{
					"X": []string{"0"},
					"Z": []string{"1"},
				})
				test(c, headers, metadata.MD{
					"X": []string{"0", "1"},
					"Y": []string{"1", "2"},
					"Z": []string{"1"},
				})
			})

			Convey("binary", func() {
				Convey("Works", func() {
					const name = "Lucy"
					b64 := base64.StdEncoding.EncodeToString([]byte(name))
					test(c, header("Name-Bin", b64), metadata.MD{
						"Name": []string{name},
					})
				})
				Convey("Fails", func() {
					c, err := parseHeader(c, header("Name-Bin", "zzz"))
					So(c, ShouldEqual, c)
					So(err, ShouldErrLike, "Name-Bin header: illegal base64 data at input byte 0")
				})
			})
		})
	})
}
