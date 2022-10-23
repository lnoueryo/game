// const protoLoader = require('@grpc/proto-loader')
import protoLoader from '@grpc/proto-loader'
import grpc from '@grpc/grpc-js'
// const grpc = require('@grpc/grpc-js')

const protoFile = '../protos/game.proto'
const options = {arrays: true}
// proto ファイルのロード
const pd = protoLoader.loadSync(protoFile, options)
// gPRC 用の動的な型定義生成
const gameProto = grpc.loadPackageDefinition(pd).proto

const target = 'localhost:50051'

export const gameClient = new gameProto.Game(target, grpc.credentials.createInsecure());

export interface Error {
    code: number;
    details: string;
}

export interface Game {
    id: number;
    name: string;
}

