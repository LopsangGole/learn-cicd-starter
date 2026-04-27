func TestCalculateSum(t *testing.T) {
    // Arrange
    inputA := 2
    inputB := 3
    expected := 5

    // Act
    actual := CalculateSum(inputA, inputB)

    // Assert
    if actual != expected {
        t.Errorf("expected %v, got %v", expected, actual)
    }
}