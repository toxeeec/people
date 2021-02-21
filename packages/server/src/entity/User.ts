import {
  BaseEntity,
  Column,
  Entity,
  JoinTable,
  ManyToMany,
  OneToMany,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { IFriendRequest } from './IFriendRequest';
import IRefreshToken from './IRefreshToken';
import IUser from './IUser';

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

  @ManyToMany('User')
  @JoinTable()
  friends: IUser[];
}
