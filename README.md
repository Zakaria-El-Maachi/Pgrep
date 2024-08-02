# Pgrep
Parallelized Grep Command (Creative Naming .. I know !)


# Concurrency Learning Project

This project was created to learn concurrency more effectively and to leverage the ease of use of goroutines to create concurrent programs in Go. I implemented various concurrency patterns such as `select`, `for`, `pipeline`, and `fan out - fan in`.

## How to run

Just Call the binary ./pgrep {-p} {pattern} {filename}\\
Adding -p option enables parallelism

## Concurrency Patterns Implemented

1. **Select**: Used for handling multiple channels in a single goroutine.
2. **Pipeline**: Implemented to process data in stages, where each stage is a set of goroutines running the same function.
3. **Fan Out - Fan In**: Utilized to distribute work across multiple goroutines and then aggregate the results.

## Performance Comparison

While the single-threaded implementation of my program significantly outperforms `grep`, the parallelized version is slower than `grep`. This performance gap could be reduced by not using the pipeline pattern, as it introduces inefficiencies like communication overhead through channels. However, the primary goal was to learn and effectively implement concurrent solutions, not necessarily to outperform `grep`.

Still, some of the changes made were the number of concurrent threads as well as using bufferized channels. In fact usin unbuffered channels is a hindrance in and of itself, as it forces the goroutines to synchronize, thus limiting concurrency ...

One of the things to not overlook is the size of the file used for benchmarking, which is 100 MB. It really is not a great benchmarking option and it limited the buffer size when reading from the file to 16 KB. The benchmark would be more telling if the file were 10 GB or larger, with a buffer size of 64 MB or so. I believe pgrep then would truly shine !

## Implementation Details

- **KMP Algorithm**: I implemented the Knuth-Morris-Pratt (KMP) string matching algorithm on the fly due to my familiarity with it from competitive programming. The KMP algorithm helps in efficiently finding all occurrences of a pattern in a text.


## Future Improvements

I honestly will probably not add these changes, as it is a relatively simple project, and these changes are really embelishments, not pertaining to concurrency, the sole reason of starting this project  in the first place.

The improvements are to add functionality, much like that of grep, where it can also fetch the line numbers instead of just fetching the byte offsets 

## Acknowledgements

- **Bash Scripts**: Most of the bash scripting was done with the help of GPT.
- **Go Code**: All of the Go code was written by yours truly, me, with love and personal care, focusing on learning and implementing concurrency effectively. (Aka no GPT)

Thank you for checking out this project. It was a valuable learning experience in understanding and implementing concurrency in Go.

