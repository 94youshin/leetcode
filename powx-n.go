func myPow(x float64, n int) float64 {
    if x >100 || x < -100 {
        return 0
    }
    if n > 1 << 31 || n < -(1<<31) {
        return 0
    }
    if n == 0 {
        return 1
    }
    if n < 0 {
        return 1 / myPow(x,-n)
    }
    if n%2 == 0 {
        return myPow(x, n/2) * myPow(x,n/2)
    }
    return myPow(x,n/2) * myPow(x,n/2) * x
}
