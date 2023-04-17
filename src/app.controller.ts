import { Bind,Param,Controller,UseInterceptors,UploadedFile, Get, Post, Put, Body } from '@nestjs/common';
import { FileInterceptor } from '@nestjs/platform-express'
import { AppService } from './app.service';
import { User,UpdateUser } from './app.dto';
import { ApiResponse } from '@nestjs/swagger';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}
  @Get('/user/:id')
  @Bind(Param())
  @ApiResponse({
    status: 200,
    description: 'get user by id',
    type: User,
  })
  getUser(params): Promise<Object> {
    return this.appService.getUser(params.id);
  }
  @Get('/users')
  @ApiResponse({
    status: 200,
    description: 'get users',
    type: class Users{
      users:User[];
      documents: 1;
    },
  })
  getUses(): Promise<Object> {
    return this.appService.getUsers();
  }

  @Post('/user')
  @UseInterceptors(FileInterceptor('file'))
  @Bind(UploadedFile(),Body())
  createUser(file,user:User):Object{
    user.picture = file
    return this.appService.createUser(user)
  }
  @Put('/user/:id')
  @UseInterceptors(FileInterceptor('file'))
  @Bind(UploadedFile(),Param(),Body())
  updateUser(file,params,user:UpdateUser):Promise<Object>{
    if (file && file.size != 0){
      user.picture = file
    }
    return this.appService.updateUser(params.id,user)
  }
}
