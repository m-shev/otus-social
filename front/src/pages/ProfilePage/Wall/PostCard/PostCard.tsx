import * as React from 'react';
import styles from './PostCard.module.scss';
import {Post} from '../../../../types';
import {useUserProfile} from '../../hooks';
import {LoadableContent, SmallDots} from '../../../../components/LoadableContent';
export type PostCardProps = {post: Post};

export const PostCard: React.FC<PostCardProps> = ({post}) => {
    const {authorId, content, updateAt, imageLink} = post;
    const {isFetch, userProfile} = useUserProfile(authorId);

    return (
        <div className={styles.root}>
            <LoadableContent isLoading={isFetch} Loader={<SmallDots />}>
                {userProfile && (
                    <div className={styles.author}>
                        <img className={styles.avatar} src={userProfile.avatar} />

                        <div className={styles.publicationInfo}>
                            <div className={styles.name}>
                                {' '}
                                {`${userProfile.name} ${userProfile.surname}`}
                            </div>

                            <div className={styles.date}>
                                {new Date(updateAt).toLocaleDateString()}
                            </div>
                        </div>
                    </div>
                )}
            </LoadableContent>

            <div className={styles.content}>{content}</div>

            <img className={styles.img} src={imageLink} alt="" />
        </div>
    );
};
