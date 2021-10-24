import * as React from 'react';
import styles from './FriendList.module.scss';
import {UserFriend} from '../../types';
import {Link} from 'react-router-dom';

export type FriendListProps = {
    list: UserFriend[];
};

export const FriendList: React.FC<FriendListProps> = ({list}) => {
    return (
        <div className={styles.root}>
            {list.map((friend) => {
                return (
                    <div className={styles.friend} key={friend.id}>
                        <img className={styles.img} src={friend.avatar} alt="" />

                        <Link className={styles.link} to={`/profile/?id=${friend.id}`}>
                            {`${friend.name} ${friend.surname}`}
                        </Link>
                    </div>
                );
            })}
        </div>
    );
};
