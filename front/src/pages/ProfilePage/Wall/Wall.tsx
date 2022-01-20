import * as React from 'react';
import styles from './Wall.module.scss';
import {usePostList} from '../hooks/usePostList';
import {PostCard} from './PostCard';

export type WallProps = {};

export const Wall: React.FC<WallProps> = () => {
    const {list} = usePostList();

    return (
        <div className={styles.root}>
            {list.map((post) => {
                return <PostCard post={post} key={post.id} />;
            })}
        </div>
    );
};
