func generateParenthesis(n int) []string {
    // end if openedN == closingN == n
    // add opened if closingN == openedN < n
    // add closing if closingN < openedN

    var (
        res = make([]string, 0)
        stack = make([]string, 0, n)
        backtrack func(openedCnt, closedCnt int)
    )
    
    backtrack = func(openedCnt, closedCnt int) {
        if openedCnt == closedCnt && closedCnt == n {
            res = append(res, strings.Join(stack, ""))
            return
        }
        
        if openedCnt < n {
            stack = append(stack, "(")
            backtrack(openedCnt + 1, closedCnt)
            stack = stack[:len(stack) - 1]
        }
        
        if closedCnt < openedCnt {
            stack = append(stack, ")")
            backtrack(openedCnt, closedCnt + 1)
            stack = stack[:len(stack) - 1]
        }
    }
    
    backtrack(0, 0)
    return res
}