package resolver

import (
	"fmt"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/interpreter/functions"
	"google.golang.org/genproto/googleapis/api/expr/v1alpha1"

	"github.com/blang/semver/v4"
)

func newCELEvaluatorProvider() *celEvaluatorProvider {
	env, err := cel.NewEnv(cel.Declarations(
		decls.NewVar("properties", decls.NewListType(decls.NewMapType(decls.String, decls.Any)))),
		cel.Lib(semverLib{}),
	)
	if err != nil {
		panic(err)
	}
	return &celEvaluatorProvider{
		env: env,
	}
}

type celEvaluatorProvider struct {
	env *cel.Env
}

type celEvaluator struct {
	p cel.Program
}

type semverLib struct{}

func (semverLib) CompileOptions() []cel.EnvOption {
	return []cel.EnvOption{
		cel.Declarations(
			decls.NewFunction("semver_compare",
				decls.NewOverload("semver_compare",
					[]*expr.Type{decls.String, decls.String},
					decls.Int))),
	}
}

func (semverLib) ProgramOptions() []cel.ProgramOption {
	return []cel.ProgramOption{
		cel.Functions(
			&functions.Overload{
				Operator: "semver_compare",
				Binary:   semverCompare,
			},
		),
	}
}

func (e celEvaluator) Evaluate(env map[string]interface{}) (bool, error) {
	result, _, err := e.p.Eval(env)
	if err != nil {
		return false, err
	}

	// we should have already ensured that this will be types.Bool during compilation
	if b, ok := result.Value().(bool); ok {
		fmt.Printf("@@@ evaluated to %t\n", b)
		return b, nil
	}

	return false, fmt.Errorf("cel expression evalutated to %T, not bool", result.Value())
}

func (e *celEvaluatorProvider) Evaluator(id, source string) (Evaluator, error) {
	if id != "cel" {
		return nil, fmt.Errorf("no available evaluator with id %q", id)
	}

	ast, issues := e.env.Compile(source)
	if err := issues.Err(); err != nil {
		return nil, err
	}
	if ast.ResultType() != decls.Bool {
		return nil, fmt.Errorf("cel expressions must have type Bool")
	}

	p, err := e.env.Program(ast)
	if err != nil {
		return nil, err
	}

	fmt.Printf("@@@ compiled %q\n", source)
	return celEvaluator{p: p}, nil
}

func semverCompare(val1, val2 ref.Val) ref.Val {
	str, ok := val1.(types.String)
	if !ok {
		return types.ValOrErr(str, "unexpected type '%v'", val1.Type())
	}

	str2, ok := val2.(types.String)
	if !ok {
		return types.ValOrErr(str, "unexpected type '%v'", val2.Type())
	}

	v1, err := semver.ParseTolerant(string(str))
	if err != nil {
		return types.ValOrErr(str, "unable to parse '%v' to semver format", val1.Value())
	}

	v2, err := semver.ParseTolerant(string(str2))
	if err != nil {
		return types.ValOrErr(str, "unable to parse '%v' to semver format", val2.Value())
	}

	return types.Int(v1.Compare(v2))
}
