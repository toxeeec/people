import {
  BaseEntity,
  Column,
  Entity,
  ManyToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';
import IUser from './IUser';

@Entity('refreshTokens')
export class RefreshToken extends BaseEntity {
  @PrimaryGeneratedColumn()
  id: number;

  @Column()
  token: string;

  @Column({ nullable: true })
  userId: string;

  @ManyToOne('User', 'tokens')
  user: IUser;
}
