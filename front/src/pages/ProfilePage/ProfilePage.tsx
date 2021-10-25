import * as React from 'react';
import {useFriendText, useInterests, useIsCurrentUser, useUserProfile} from './hooks';
import {LoadableContent} from '../../components/LoadableContent';
import styles from './ProfilePage.module.scss';
import {UserGender} from '../../types';
import {Header} from '../../components/Header';
import {ReferenceAction} from './ReferenceAction';
import {FriendList} from '../../components/FriendList';
import {Link} from 'react-router-dom';

export const ProfilePage: React.FC = () => {
    const {userProfile, isFetch, loadProfile} = useUserProfile();
    const isCurrentUser = useIsCurrentUser(userProfile);
    const interests = useInterests(userProfile?.interests || []);
    const gender = userProfile?.gender === UserGender.Female ? 'женский' : 'мужской';
    const friendText = useFriendText(userProfile, isCurrentUser);

    return (
        <div className={styles.root}>
            <Header />

            <LoadableContent isLoading={isFetch}>
                {userProfile && (
                    <>
                        <div className={styles.user}>
                            <div className={styles.leftColumn}>
                                <img className={styles.img} src={userProfile.avatar} alt="" />

                                <ReferenceAction
                                    userProfile={userProfile}
                                    loadProfile={loadProfile}
                                />
                            </div>

                            <div className={styles.rightColumn}>
                                <div className={styles.text}>
                                    {`${userProfile.name} ${userProfile.surname}`}
                                </div>

                                <div>{`Возраст ${userProfile.age} лет`}</div>

                                <div>{`Пол ${gender}`}</div>

                                <div>{`Город ${userProfile.city}`}</div>
                                {interests && (
                                    <>
                                        <div>{`Интересы: `}</div>
                                        <div>{interests}</div>
                                    </>
                                )}
                            </div>
                        </div>

                        {isCurrentUser && (
                            <Link to={'/find'} className={styles.link}>
                                Найти друзей
                            </Link>
                        )}

                        <div className={styles.friend}>{friendText}</div>

                        <div className={styles.friendList}>
                            <FriendList list={userProfile.friends} />
                        </div>
                    </>
                )}
            </LoadableContent>
        </div>
    );
};
