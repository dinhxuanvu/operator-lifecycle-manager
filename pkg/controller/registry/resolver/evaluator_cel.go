package resolver

import (
	"fmt"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func newCELEvaluatorProvider() *celEvaluatorProvider {
	env, err := cel.NewEnv(cel.Declarations(
		decls.NewVar("properties", decls.NewListType(decls.NewMapType(decls.String, decls.Any))),
	))
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
