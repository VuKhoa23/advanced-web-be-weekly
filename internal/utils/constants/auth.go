package constants

import "time"

const JWT_DURATION = 2 * time.Minute
const REFRESH_TOKEN_DURATION = 30 * 24 * time.Hour // 30 days
const COOKIE_DURATION = 2629800 // 1 month

// const JWT_DURATION = 5 * time.Second
// const REFRESH_TOKEN_DURATION = 20 * time.Second
// const COOKIE_DURATION = 2629800
