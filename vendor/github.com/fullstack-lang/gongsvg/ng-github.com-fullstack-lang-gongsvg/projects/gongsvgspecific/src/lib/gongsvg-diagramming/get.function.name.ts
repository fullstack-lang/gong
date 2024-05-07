export function getFunctionName() {
    // Create a new error and split its stack trace
    const stack = new Error().stack;
    if (!stack) return '';

    // Split the stack trace and find the caller
    const stackLines = stack.split('\n');
    // console.log("stack", stack)
    const callerIndex = 2; // This might change depending on the browser/environment
    if (stackLines.length < callerIndex + 1) return '';

    // Extract the function name
    const callerLine = stackLines[callerIndex];
    const match = callerLine.match(/at (\S+)/);
    return match ? match[1] : '';
}