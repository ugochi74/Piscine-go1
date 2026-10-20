def parse_number_safely(user_input):
    """
    Parse a number from a string, handling common formatting issues.
    Returns a tuple: (parsed_value, error_message)
    """
    # Check for empty input
    if not user_input or user_input.strip() == "":
        return None, "Empty input provided."
    # Clean the input: remove spaces and commas
    cleaned = user_input.replace(",", "").strip()
    cleaned = cleaned.replace("₦", "")
    # Try to convert to float first (handles integers too)
    try:
        return float(cleaned), None
    except ValueError:
        return None, f"'{user_input}' is not a valid number."
# Test the function
test_inputs = ["42", "3.14", "2,500.50", "₦1500", "hello", "", "  100  "]
for inp in test_inputs:
    value, error = parse_number_safely(inp)
    if error:
        print(f"❌ Input: '{inp}' -> Error: {error}")
    else:
        print(f"✅ Input: '{inp}' -> Value: {value}")
