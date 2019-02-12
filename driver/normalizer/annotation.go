package normalizer

import . "gopkg.in/bblfsh/sdk.v2/uast/transformer"

var Native = Transformers([][]Transformer{}...)

var Code = []CodeTransformer{}

var Annotations = []Mapping{}
