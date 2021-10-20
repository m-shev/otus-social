import {useCallback} from 'react';
import {CreateUserForm, HttpStatus, IUseRequestState} from '../../../types';
import {createUserPost} from '../../../api';
import {CreateFormValues} from '../const';
import {FormikHelpers} from 'formik/dist/types';
import {useRequest} from '../../../hooks';

export interface IUseCreateUser extends IUseRequestState {
    onSubmit: (values: CreateFormValues, formikHelpers: FormikHelpers<CreateFormValues>) => void;
}

const PASSWORDS_SHOULD_BE_EQUAL = 'Пароли должны совпадать';

const mapValuesToForm = (values: CreateFormValues): CreateUserForm => {
    const {firstName, lastName, password2, gender, age, ...rest} = values;

    return {
        name: firstName,
        surname: lastName,
        age: age as number,
        ...rest,
    };
};

export const useCreateUser = (): IUseCreateUser => {
    const {error, setError, isFetch, setIsFetch, setRequestState, requestState} = useRequest();

    const onSubmit = useCallback(
        async (
            values: CreateFormValues,
            helpers: FormikHelpers<CreateFormValues>,
        ): Promise<void> => {
            if (values.password !== values.password2) {
                helpers.setErrors({
                    password: PASSWORDS_SHOULD_BE_EQUAL,
                    password2: PASSWORDS_SHOULD_BE_EQUAL,
                });
                return;
            }

            setIsFetch(true);

            const resp = await createUserPost(mapValuesToForm(values));

            if (resp.status === HttpStatus.Ok) {
                setRequestState('success');
            } else {
                setRequestState('fail');
                const errorText = await resp.text();
                setError(new Error(errorText));
            }
            setIsFetch(false);
        },
        [setError, setIsFetch, setRequestState],
    );

    return {
        isFetch,
        requestState,
        error,
        onSubmit,
    };
};
