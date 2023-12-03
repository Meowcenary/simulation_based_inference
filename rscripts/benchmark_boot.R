# Choose a CRAN mirror (replace 'mirror_url' with the URL of a CRAN mirror near you)
mirror_url <- "https://cran.rstudio.com/"
chooseCRANmirror(ind = 1)  # You can choose a mirror interactively or set it explicitly
options(repos = structure(c(CRAN = mirror_url)))

# Check if 'boot' package is installed and if not install it
if (!require("boot", character.only = TRUE)) {
  install.packages("boot")
}

# Load 'boot' package
library(boot)

# Create a function to calculate the mean
mean_function <- function(data, indices) {
  sample_data <- data[indices]
  return(mean(sample_data))
}

benchmark_vector <- function() {
  # Start timer for benchmarking
  start_time <- Sys.time()

  # Generate a random dataset
  set.seed(123)  # for reproducibility
  data_vector <- rnorm(1000)
  result_vector <- boot(data_vector, mean_function, R = 1000)


  # End timer for benchmarking
  end_time <- Sys.time()
  print("Benchmark vector")
  # Print difference end to start
  print(end_time - start_time)
}

benchmark_matrix <- function() {
  # Start timer for benchmarking
  start_time <- Sys.time()

  # Generate a random dataset
  set.seed(123)  # for reproducibility
  rows = 10000
  cols = 20
  data_matrix <- matrix(rnorm(rows * cols), nrow = rows, ncol = cols)
  result_matrix <- boot(data_matrix, mean_function, R = 1000)

  # End timer for benchmarking
  end_time <- Sys.time()
  # Print difference end to start
  print("Benchmark matrix-")
  print(end_time - start_time)
}

benchmark_vector()
benchmark_matrix()
