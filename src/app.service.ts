import { Injectable, OnModuleInit } from '@nestjs/common';
import { UpdateUser, User } from './app.dto';
import { ConfigService } from '@nestjs/config';
import {  connect } from 'mongoose';
import { createLogger,format,transports } from 'winston';

let logger;

@Injectable()
export class AppService {
  constructor(private configService: ConfigService) {}
  onModuleInit() {
    connect('mongodb://'+this.configService.get<string>('database.host')+':'+this.configService.get<number>('database.port')+'/testDB')
    const env = this.configService.get<string>('env')
    if (env == 'dev'){
      logger = createLogger({
        format: format.combine(format.timestamp(), format.json()),
        defaultMeta: {
          service: "js-user-admin",
        },
        transports:[new transports.File({filename: "/tmp/logs/microservice.log"})],
      });
      return
    }
    logger = createLogger({
      format: format.combine(format.timestamp(), format.json()),
      defaultMeta: {
        service: "js-user-admin",
      },
      transports:[ new transports.Console()],
    });
    
  }
  async getUser(id:string): Promise<User> {
    let u = new User(logger=logger)
    return u.getUser(id)
  }

  getUsers(): Promise<Object> {
    let u = new User(logger=logger)
    return u.getUsers();
  }

  createUser(user:User):Object{
    let u = new User(logger,user.name,user.last_name,user.address,user.picture)
    return u.newUser()
  }
  async updateUser(id:string,user:UpdateUser):Promise<Object>{
    let u = new UpdateUser(logger=logger,user.name,user.last_name,user.address,user.picture)
    return u.updateUser(id)
  }
}
