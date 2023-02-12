function foo(){
  console.log("foo");
}

function bar(){
  console.log("bar");
}

for (var i = 0; i < 10; i++){
  if (i%2 === 0){
    foo();
  } else {
    bar();
  }
}
console.log("done");

