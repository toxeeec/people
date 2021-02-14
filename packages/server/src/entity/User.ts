import {
  BaseEntity,
  Column,
  Entity,
  OneToMany,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { IFriendRequest } from './IFriendRequest';
import IRefreshToken from './IRefreshToken';

@Entity('users')
export class User extends BaseEntity {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  @Column()
  name: string;

  @Column()
  surname: string;

  @Column()
  email: string;

  @Column()
  password: string;

  @OneToMany('RefreshToken', 'user')
  tokens: IRefreshToken[];

  @OneToMany('FriendRequest', 'sender')
  friendRequestsSent: IFriendRequest[];

  @OneToMany('RefreshToken', 'receiver')
  friendRequestsReceiver: IFriendRequest[];
}
