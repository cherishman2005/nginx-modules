function sleep(ms){
  return new Promise((resolve,reject) => setTimeout(
     ()=>resolve(),ms
  ))
}

async function hello(){
  await sleep(3000)
  console.log("hello")
};

async function test(){
  await sleep(3000)
  console.log("world")
};

await hello();
await test();