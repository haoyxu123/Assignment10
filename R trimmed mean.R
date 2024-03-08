set.seed(123)

float_data <- runif(100, min = 0, max = 100)

int_data <- sample(100, replace = TRUE)

float_trimmed_mean <- mean(float_data, trim = 0.05)
int_trimmed_mean <- mean(int_data, trim = 0.05)

# Print the results
cat("Trimmed mean of float data in R:", float_trimmed_mean, "\n")
cat("Trimmed mean of integer data in R:", int_trimmed_mean, "\n")
