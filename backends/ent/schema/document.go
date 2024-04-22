// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileName: ent/schema/document.go
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// ------------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ------------------------------------------------------------------------
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"
)

type Document struct {
	ent.Schema
}

func (Document) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SourceDataMixin{},
	}
}

func (Document) Fields() []ent.Field { return nil }

func (Document) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metadata", Metadata.Type).Ref("document").Required().Unique(),
		edge.From("node_list", NodeList.Type).Ref("document").Required().Unique(),
	}
}

func (Document) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("metadata", "node_list").Unique(),
	}
}

func (Document) Annotations() []schema.Annotation { return nil }
