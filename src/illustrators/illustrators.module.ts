import { Module } from '@nestjs/common';
import { IllustratorsService } from './illustrators.service';
import { IllustratorsController } from './illustrators.controller';

@Module({
  controllers: [IllustratorsController],
  providers: [IllustratorsService],
})
export class IllustratorsModule {}
