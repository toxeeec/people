import {
  BaseEntity,
  Column,
  Entity,
  ManyToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';
import IUser from './IUser';

export enum FriendRequestStatus {
  PENDING,
  ACCEPTED,
  BLOCKED,
}

@Entity('friendRequests')
export class FriendRequest extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @ManyToOne('User', 'friendRequestsSent')
  sender: IUser;

  @ManyToOne('User', 'friendRequestsReceived')
  receiver: IUser;

  @Column({ default: FriendRequestStatus.PENDING })
  status: FriendRequestStatus;
}
