# Copyright 2018 The LUCI Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This module is not imported by anything. It exists only to document public
# native symbols exposed by `lucicfg` go code in the global namespace of all
# modules.


def _to_textpb(msg):
  """Serializes a protobuf message to a string using ASCII proto serialization.

  Args:
    msg: a proto message to serialize. Required.
  """


def _to_jsonpb(msg):
  """Serializes a protobuf message to a string using JSONPB serialization.

  Args:
    msg: a proto message to serialize. Required.
  """


def _from_textpb(ctor, text):
  """Deserializes a protobuf message given its ASCII proto serialization.

  Args:
    ctor: a message constructor function, same one you would normally use to
        create a new message. Required.
    text: a string with the serialized message. Required.

  Returns:
    Deserialized message constructed via `ctor`.
  """


def _from_jsonpb(ctor, text):
  """Deserializes a protobuf message given its JSONPB serialization.

  Args:
    ctor: a message constructor function, same one you would normally use to
        create a new message. Required.
    text: a string with the serialized message. Required.

  Returns:
    Deserialized message constructed via `ctor`.
  """

def _struct_to_textpb(s):
  """Converts a struct to a text proto string.

  Args:
    s: a struct object. May not contain dicts.

  Returns:
    A str containing a text format protocol buffer message.
  """

proto = struct(
    to_textpb = _to_textpb,
    to_jsonpb = _to_jsonpb,
    from_textpb = _from_textpb,
    from_jsonpb = _from_jsonpb,
    struct_to_textpb = _struct_to_textpb,
)
