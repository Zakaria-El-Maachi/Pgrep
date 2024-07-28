# Pgrep
Parallelized Grep Command (Creative Naming .. I know !)


# Concurrency Learning Project

This project was created to learn concurrency more effectively and to leverage the ease of use of goroutines to create concurrent programs in Go. I implemented various concurrency patterns such as `select`, `for`, `pipeline`, and `fan out - fan in`.

## Concurrency Patterns Implemented

1. **Select**: Used for handling multiple channels in a single goroutine.
2. **Pipeline**: Implemented to process data in stages, where each stage is a set of goroutines running the same function.
3. **Fan Out - Fan In**: Utilized to distribute work across multiple goroutines and then aggregate the results.

## Performance Comparison

While the single-threaded implementation of my program significantly outperforms `grep`, the parallelized version is slower than `grep`. This performance gap could be reduced by not using the pipeline pattern, as it introduces inefficiencies like communication overhead through channels. However, the primary goal was to learn and effectively implement concurrent solutions, not necessarily to outperform `grep`.

## Implementation Details

- **KMP Algorithm**: I implemented the Knuth-Morris-Pratt (KMP) string matching algorithm on the fly due to my familiarity with it from competitive programming. The KMP algorithm helps in efficiently finding all occurrences of a pattern in a text.

## Acknowledgements

- **Bash Scripts**: Most of the bash scripting was done with the help of GPT.
- **Go Code**: All of the Go code was written with love and personal care, focusing on learning and implementing concurrency effectively.

Thank you for checking out this project. It was a valuable learning experience in understanding and implementing concurrency in Go.

