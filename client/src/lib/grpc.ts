// const protoLoader = require('@grpc/proto-loader')
import protoLoader from '@grpc/proto-loader'
import grpc from '@grpc/grpc-js'
// const grpc = require('@grpc/grpc-js')

const protoFile = '../protos/table.proto'
const options = {arrays: true}
// proto ファイルのロード
const pd = protoLoader.loadSync(protoFile, options)
// gPRC 用の動的な型定義生成
const hello_proto = grpc.loadPackageDefinition(pd).table

const target = 'localhost:50051'

export const gameClient = new hello_proto.Game(target, grpc.credentials.createInsecure());
