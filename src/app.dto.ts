import { IsNotEmpty } from 'class-validator';
import { Schema,Types, model } from 'mongoose';
import { Logger } from 'winston';

interface IUser {
    name: string;
    last_name: string;
    address: string;
    picture: File;
}
const userSchema = new Schema<IUser>({
    name: { type: String, required: true },
    last_name: { type: String, required: true },
    address: { type: String, required: true },
    picture: Object
});
const UserModel = model<IUser>('User',userSchema)

export class User {
    private logger:Logger;
    id: string;
    @IsNotEmpty()
    name: string;
    last_name: string;
    address: string;
    picture: File;
    constructor(logger?,name?,last_name?,address?,picture?){
        if (this.logger == undefined && logger != undefined) {
            this.logger = logger
        }
        if (name != undefined){
            this.name = name
        }
        if (last_name != undefined){
            this.last_name = last_name
        }
        if (address != undefined){
            this.address = address
        }
        if (picture != undefined){
            this.picture = picture
        }
    }
    newUser():Object {
        this.logger.info("[ \"function\": \"newUser\" ] [ \"full_name\": \""+this.last_name+" , "+this.name+"\" ]")
        const u = new UserModel({
            name: this.name,
            last_name: this.last_name,
            address: this.address,
            picture: this.picture
        });
        u.save()
        return {'id':u._id, 'documents':1}
    };
    getUser(id:string):Promise<User>{
        this.logger.info("[ \"function\": \"getUser\" ] [ id: \""+id+"\" ]")
        const u =  UserModel.findOne({ _id:new Types.ObjectId(id)}).exec().then(u => {
            let user = new User();
            if (u) {
                user= new User(null,u.name,u.last_name,u.address,u.picture)
                this.logger.info("[ \"id\": \""+id+"\" ] [ \"full_name\": \""+u.last_name+" , "+u.name+"\" ] [ \"status\": \"ok\" ]")
            }else{
                this.logger.error("[ \"id\": \""+id+"\" ] [ \"status\": \"user not found\" ]")
            }
            user.id = id
            return user
        })
        return u
    };
    getUsers():Promise<Object>{
        this.logger.info("[ \"function\": \"getUsers\" ]")
        let result = []
        const u =  UserModel.find().exec().then(users => {
            for(var i in users) {
                const subU = users[i]
                let user = new User(null,subU.name,subU.last_name,subU.address,subU.picture)
                user.id = subU._id.toString()
                result.push(user)
             }
            return {"users":result,"documents":result.length}
        })
        return u
    }
}

export class UpdateUser {
    private logger?:Logger;
    id?: string;
    name?: string;
    last_name?: string;
    address?: string;
    picture?: File;
    constructor(logger?,name?,last_name?,address?,picture?){
        if (this.logger == undefined && logger != undefined) {
            this.logger = logger
        }
        if (name != undefined){
            this.name = name
        }
        if (last_name != undefined){
            this.last_name = last_name
        }
        if (address != undefined){
            this.address = address
        }
        if (picture != undefined){
            this.picture = picture
        }
    }
    updateUser(id:string):Promise<Object>{
        this.logger.info("[ function: \"updateUser\" ] [ id: \""+id+"\" ]")
        let userL = new UpdateUser(null,this.name,this.last_name,this.address,this.picture)
        const u = UserModel.findOneAndUpdate( {_id:new Types.ObjectId(id)},userL).exec().then( r => {
            let result;
            userL.id = id
            if (r) {
                this.logger.info("[ function: \"updateUser\" ] [ id: \""+id+"\" ] [ status: \"ok\" ]")
                result = {'users':[userL], 'documents': 1}
            }else{
                this.logger.error("[ function: \"updateUser\" ] [ id: \""+id+"\" ] [ status: \"user not found\" ]")
                result = {'users':[userL], 'documents': 0}
            }
            return result
        });
        return u
    };
}