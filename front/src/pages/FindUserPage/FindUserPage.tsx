import * as React from 'react';
import styles from './FindUserPage.module.scss';
import {fieldList, initialValues} from './const';
import {useFormik} from 'formik';
import {Field} from '../../components/Field';
import {Header} from '../../components/Header';
import {LoadableContent, SmallDots} from '../../components/LoadableContent';
import {useLastUserList, useOnFind} from './hooks';
import {FriendList} from '../../components/FriendList';

export const FindUserPage: React.FC = () => {
    const userList = useLastUserList();
    const {findList, isFetch, onSubmit, firstSearchFlag} = useOnFind();
    const formik = useFormik({
        initialValues,
        onSubmit,
    });

    const searchText = findList.length ? 'Посмотрите кого мы нашли' : 'Упс, мы никого не нашли';

    return (
        <div className={styles.root}>
            <Header />

            <h2>Поиск друзей</h2>

            <form onSubmit={formik.handleSubmit} className={styles.form}>
                {fieldList.map((field) => {
                    return (
                        <Field
                            key={field.id}
                            error={formik.errors[field.id]}
                            setFieldValue={formik.setFieldValue}
                            {...formik.getFieldProps(field.id)}
                            {...field}
                        />
                    );
                })}

                <LoadableContent isLoading={isFetch} Loader={<SmallDots />}>
                    <button type="submit" className={styles.btn}>
                        Поиск
                    </button>
                </LoadableContent>
            </form>

            {!firstSearchFlag && (
                <>
                    <div className={styles.text}>Недавно зарегистрировались</div>

                    <div className={styles.friendList}>
                        <FriendList list={userList} />
                    </div>
                </>
            )}

            {firstSearchFlag && (
                <>
                    <div className={styles.text}>{searchText}</div>

                    <div className={styles.friendList}>
                        <FriendList list={findList} />
                    </div>
                </>
            )}
        </div>
    );
};
