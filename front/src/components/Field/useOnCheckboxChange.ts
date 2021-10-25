import {SetFieldValue} from './types';
import {SyntheticEvent, useCallback} from 'react';

export const useOnCheckboxChange = (
    setFieldValue: SetFieldValue,
    fieldId: string,
    value: unknown,
): ((e: SyntheticEvent<HTMLInputElement>) => void) => {
    return useCallback(
        (e: SyntheticEvent<HTMLInputElement>) => {
            const newValue = Array.isArray(value) ? [...value] : [];

            if (e.currentTarget.checked) {
                newValue.push(e.currentTarget.value);
                setFieldValue(fieldId, newValue);
            } else {
                setFieldValue(
                    fieldId,
                    newValue.filter((item) => {
                        return item === e.currentTarget.value;
                    }),
                );
            }
        },
        [fieldId, setFieldValue, value],
    );
};
