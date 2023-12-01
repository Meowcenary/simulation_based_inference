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

# Generate a random dataset
set.seed(123)  # for reproducibility

### Matrix

# Using rnorm function for normally distributed values
rows <- 100
cols <- 1

# Generate a matrix of floats with a mean of 0 and standard deviation of 1
# data <- matrix(rnorm(rows * cols), nrow = rows, ncol = cols)

### Vector
data <- rnorm(100)

print(data)

# Create a function to calculate the mean
mean_function <- function(data, indices) {
  sample_data <- data[indices]
  return(mean(sample_data))
}

# Use the boot function for bootstrapping
result <- boot(data, mean_function, R = 1000)

# Display the bootstrap results
print(result)

# Calculate and print the bootstrap confidence interval for the mean
boot_ci <- boot.ci(result, type = "basic")
# print(boot_ci)
# Access the call element to see the original function call
# print(boot_ci$call)
