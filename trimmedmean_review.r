# Set the base folder path where your CSVs are
file_path <- "C:/Users/jpetk/Documents/_Data Science/Classes/2025 - 2026 Academic Year/MSDS 431/Module 9_1/"

# --- Example 1: Integers CSV ---
int_file <- paste0(file_path, "int_sample.csv")
x_int <- scan(int_file, sep = ",")  # reads the integers

# Symmetric trimmed mean (5%)
cat("Symmetric trimmed mean (integers):", mean(x_int, trim = 0.05), "\n")

# Asymmetric trimmed mean (low=5%, high=30%)
asym_trimmed_mean <- function(x, low = 0.05, high = 0.30) {
  low_q <- quantile(x, low)
  high_q <- quantile(x, 1 - high)
  x_trimmed <- x[x >= low_q & x <= high_q]
  mean(x_trimmed)
}
cat("Asymmetric trimmed mean (integers):", asym_trimmed_mean(x_int, low = 0.05, high = 0.30), "\n")

# --- Example 2: Floats CSV ---
float_file <- paste0(file_path, "float_sample.csv")
x_float <- scan(float_file, sep = ",")  # reads the floats

# Symmetric trimmed mean (5%)
cat("Symmetric trimmed mean (floats):", mean(x_float, trim = 0.05), "\n")

# Asymmetric trimmed mean (low=5%, high=20%)
cat("Asymmetric trimmed mean (floats):", asym_trimmed_mean(x_float, low = 0.05, high = 0.20), "\n")