import { BaseEntity, Entity, ManyToOne, PrimaryGeneratedColumn } from 'typeorm';
import IUser from './IUser';

@Entity('friendRequests')
export class FriendRequest extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @ManyToOne('User', 'friendRequestsSent')
  sender: IUser;

  @ManyToOne('User', 'friendRequestsReceived')
  receiver: IUser;
}
