// A simple test program to make sure concolic execution is working

function C$symbol() {}

let x = C$symbol("X", "Real");

if (x > 42) {
    throw 'Crash? ... Yeah, burn? ... Make a wish.';
}
