// const protoLoader = require('@grpc/proto-loader')
import protoLoader from '@grpc/proto-loader'
import grpc from '@grpc/grpc-js'
// const grpc = require('@grpc/grpc-js')

const protoFile = '../protos/helloworld.proto'
// proto ファイルのロード
const pd = protoLoader.loadSync(protoFile)
// gPRC 用の動的な型定義生成
const hello_proto = grpc.loadPackageDefinition(pd).helloworld

const target = 'localhost:50051'
// console.log(new hello_proto.Game(target, grpc.credentials.createInsecure()).GetTables)
// export const greeterClient = new hello_proto.Greeter(target, grpc.credentials.createInsecure());
export const gameClient = new hello_proto.Game(target, grpc.credentials.createInsecure());
// gameClient.GetTables({id: 1}, (err, response) => {
//     console.log(err)
//     console.log('practice:', response);
//   })
  // greeterClient.SayHello({ name: 'user' }, (err, response) => {
  //   console.log('Greeting:', response);
  // });
// console.log(new hello_proto.Greeter(target, grpc.credentials.createInsecure()).sayHello)
// console.log(gameClient.getTables({id: 1}, (err, res)=>{return res}))