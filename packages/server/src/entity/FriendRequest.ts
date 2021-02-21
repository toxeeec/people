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

  @Column({ nullable: true })
  senderId: string;

  @ManyToOne('User', 'friendRequestsSent')
  sender: IUser;

  @Column({ nullable: true })
  receiverId: string;

  @ManyToOne('User', 'friendRequestsReceived')
  receiver: IUser;

  @Column({ default: FriendRequestStatus.PENDING })
  status: FriendRequestStatus;
}
