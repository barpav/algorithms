package basic_calculator_2

func calculate(s string) int {
	exp := expression{}
	var i integer
	for _, char := range []byte(s) {
		value := char - '0'
		switch {
		case value < 10:
			i = i*10 + integer(value)
		case char == '+' || char == '-' || char == '*' || char == '/':
			exp.proceed(i, char)
			i = 0
		}
	}
	exp.proceed(i, ';')
	return int(exp.value())
}

type expressionMember interface {
	value() integer
}

type expression struct {
	x, y     *expressionMember
	operator byte
}

func (e *expression) value() integer {
	x := *(e.x)
	if e.y == nil {
		return x.value()
	}
	y := *(e.y)
	operation := operation(e.operator)
	return operation(x.value(), y.value())
}

func (e *expression) proceed(i integer, operator byte) {
	if e.x == nil {
		var x expressionMember = &i
		e.x, e.operator = &x, operator
		return
	}

	if e.y == nil {
		if (e.operator == '+' || e.operator == '-') && (operator == '*' || operator == '/') {
			yExp := expression{}
			var x expressionMember = &i
			yExp.x, yExp.operator = &x, operator
			var subExp expressionMember = &yExp
			e.y = &subExp
		} else {
			var v expressionMember = &i
			e.y = &v
			if operator != ';' {
				value := e.value()
				v = &value
				e.x, e.operator, e.y = &v, operator, nil
			}
		}
	} else {
		y := *(e.y)
		var yExp *expression
		yExp, y = y.(*expression), &i
		yExp.y = &y
		value := yExp.value()
		var v expressionMember = &value
		if yExp.operationHasChanged(operator) || operator == ';' {
			e.y = &v
			if operator == ';' {
				return
			}
			value = e.value()
			v = &value
			e.x, e.operator, e.y = &v, operator, nil
		} else {
			yExp.x, yExp.operator, yExp.y = &v, operator, nil
		}
	}
}

func (e *expression) operationHasChanged(newOperator byte) bool {
	var hasChanged bool
	switch {
	case e.operator == '+' || e.operator == '-':
		hasChanged = newOperator == '*' || newOperator == '/'
	case e.operator == '*' || e.operator == '/':
		hasChanged = newOperator == '+' || newOperator == '-'
	}
	return hasChanged
}

type integer int

func (i *integer) value() integer {
	return *i
}

type expressionOperation func(x, y integer) integer

func operation(operator byte) expressionOperation {
	switch operator {
	case '+':
		return add
	case '-':
		return sub
	case '*':
		return mul
	case '/':
		return div
	}
	return nil
}

func add(x, y integer) integer {
	return x + y
}

func sub(x, y integer) integer {
	return x - y
}

func mul(x, y integer) integer {
	return x * y
}

func div(x, y integer) integer {
	return x / y
}
