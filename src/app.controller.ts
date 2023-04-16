import { Bind,Param,Controller,UseInterceptors,UploadedFile, Get, Post, Put, Body } from '@nestjs/common';
import { FileInterceptor } from '@nestjs/platform-express'
import { AppService } from './app.service';
import { User,UpdateUser } from './app.dto';

@Controller()
export class AppController {
  constructor(private readonly appService: AppService) {}
  @Get('/user/:id')
  @Bind(Param())
  getUser(params): Promise<Object> {
    return this.appService.getUser(params.id);
  }
  @Get('/users')
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
    if (file.size != 0){
      user.picture = file
    }
    return this.appService.updateUser(params.id,user)
  }
}
