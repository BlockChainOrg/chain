package txvm

//go:generate sh gen.sh

var OpNames = [...]string{
	Fail:              "fail",
	PC:                "pc",
	Exec:              "exec",
	JumpIf:            "jumpif",
	Roll:              "roll",
	Bury:              "bury",
	Depth:             "depth",
	Drop:              "drop",
	Dup:               "dup",
	ToAlt:             "toalt",
	FromAlt:           "fromalt",
	Len:               "len",
	Field:             "field",
	Equal:             "equal",
	Not:               "not",
	And:               "and",
	Or:                "or",
	GT:                "gt",
	GE:                "ge",
	Abs:               "abs",
	Add:               "add",
	Mul:               "mul",
	Div:               "div",
	Mod:               "mod",
	Lshift:            "lshift",
	Rshift:            "rshift",
	Cat:               "cat",
	Slice:             "slice",
	BitNot:            "bitnot",
	BitAnd:            "bitand",
	BitOr:             "bitor",
	BitXor:            "bitxor",
	SHA256:            "sha256",
	SHA3:              "sha3",
	CheckSig:          "checksig",
	CheckMultiSig:     "checkmultisig",
	PointAdd:          "pointadd",
	PointSub:          "pointsub",
	PointMul:          "pointmul",
	Cond:              "cond",
	Unlock:            "unlock",
	UnlockOutput:      "unlockoutput",
	Merge:             "merge",
	Split:             "split",
	ProveRange:        "proverange",
	ProveValue:        "provevalue",
	ProveAsset:        "proveasset",
	Blind:             "blind",
	Lock:              "lock",
	Satisfy:           "satisfy",
	Anchor:            "anchor",
	Issue:             "issue",
	IssueCA:           "issueca",
	Retire:            "retire",
	Varint:            "varint",
	Encode:            "encode",
	VM1CheckPredicate: "vm1checkpredicate",
	VM1Unlock:         "vm1unlock",
	VM1Nonce:          "vm1nonce",
	VM1Issue:          "vm1issue",
	VM1Mux:            "vm1mux",
	VM1Withdraw:       "vm1withdraw",
	Nop0:              "nop0",
	Nop1:              "nop1",
	Nop2:              "nop2",
	Nop3:              "nop3",
	Nop4:              "nop4",
	Nop5:              "nop5",
	Nop6:              "nop6",
	Nop7:              "nop7",
	Nop8:              "nop8",
	Private:           "private",
}

var OpCodes = map[string]byte{
	"fail":              Fail,
	"pc":                PC,
	"exec":              Exec,
	"jumpif":            JumpIf,
	"roll":              Roll,
	"bury":              Bury,
	"depth":             Depth,
	"drop":              Drop,
	"dup":               Dup,
	"toalt":             ToAlt,
	"fromalt":           FromAlt,
	"len":               Len,
	"field":             Field,
	"equal":             Equal,
	"not":               Not,
	"and":               And,
	"or":                Or,
	"gt":                GT,
	"ge":                GE,
	"abs":               Abs,
	"add":               Add,
	"mul":               Mul,
	"div":               Div,
	"mod":               Mod,
	"lshift":            Lshift,
	"rshift":            Rshift,
	"cat":               Cat,
	"slice":             Slice,
	"bitnot":            BitNot,
	"bitand":            BitAnd,
	"bitor":             BitOr,
	"bitxor":            BitXor,
	"sha256":            SHA256,
	"sha3":              SHA3,
	"checksig":          CheckSig,
	"checkmultisig":     CheckMultiSig,
	"pointadd":          PointAdd,
	"pointsub":          PointSub,
	"pointmul":          PointMul,
	"cond":              Cond,
	"unlock":            Unlock,
	"unlockoutput":      UnlockOutput,
	"merge":             Merge,
	"split":             Split,
	"proverange":        ProveRange,
	"provevalue":        ProveValue,
	"proveasset":        ProveAsset,
	"blind":             Blind,
	"lock":              Lock,
	"satisfy":           Satisfy,
	"anchor":            Anchor,
	"issue":             Issue,
	"issueca":           IssueCA,
	"retire":            Retire,
	"varint":            Varint,
	"encode":            Encode,
	"vm1checkpredicate": VM1CheckPredicate,
	"vm1unlock":         VM1Unlock,
	"vm1nonce":          VM1Nonce,
	"vm1issue":          VM1Issue,
	"vm1mux":            VM1Mux,
	"vm1withdraw":       VM1Withdraw,
	"nop0":              Nop0,
	"nop1":              Nop1,
	"nop2":              Nop2,
	"nop3":              Nop3,
	"nop4":              Nop4,
	"nop5":              Nop5,
	"nop6":              Nop6,
	"nop7":              Nop7,
	"nop8":              Nop8,
	"private":           Private,
}
