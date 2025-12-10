package work

import "runtime"

// EnableLocksProfiling включает профили блокировок и мьютексов.
// Это доп. задание, можно вызвать из main при необходимости.
func EnableLocksProfiling() {
    runtime.SetBlockProfileRate(1)     // включить Block profile
    runtime.SetMutexProfileFraction(1) // включить Mutex profile
}
