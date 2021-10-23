import * as React from 'react';
import {useInterests, useUserProfile} from './hooks';
import {LoadableContent} from '../../components/LoadableContent';
import styles from './ProfilePage.module.scss';
import {UserGender} from '../../types';
import {Header} from '../../components/Header';
import {ReferenceAction} from './ReferenceAction';

export type UserPageProps = {};

export const ProfilePage: React.FC<UserPageProps> = () => {
    const {userProfile, isFetch} = useUserProfile();

    const interests = useInterests(userProfile?.interests || []);
    const gender = userProfile?.gender === UserGender.Female ? 'женский' : 'мужской';
    return (
        <div className={styles.root}>
            <Header />

            <LoadableContent isLoading={isFetch}>
                {userProfile && (
                    <div className={styles.user}>
                        <div className={styles.leftColumn}>
                            <img className={styles.img} src={userProfile.avatar} alt="" />

                            <ReferenceAction userProfile={userProfile} />
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
                )}
            </LoadableContent>
        </div>
    );
};
