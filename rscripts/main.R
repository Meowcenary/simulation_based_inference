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

# Create a function to calculate the mean
mean_function <- function(data, indices) {
  sample_data <- data[indices]
  return(mean(sample_data))
}

### Matrix data
# Using rnorm function for normally distributed values
rows <- 100
cols <- 2
# Generate a matrix of floats with a mean of 0 and standard deviation of 1
dataMatrix <- matrix(rnorm(rows * cols), nrow = rows, ncol = cols)
resultMatrix <- boot(dataMatrix, mean_function, R = 1000)
# Display the matrix results
print(resultMatrix)

### Vector data
dataVector <- rnorm(100)
# Use the boot function for bootstrapping
resultVector <- boot(dataVector, mean_function, R = 1000)
# Display the bootstrap results
print(resultVector)
