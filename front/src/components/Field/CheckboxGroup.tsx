import * as React from 'react';
import {CheckboxFieldProps} from './types';
import styles from './Field.module.scss';
import {useOnCheckboxChange} from './useOnCheckboxChange';

export type CheckboxProps = CheckboxFieldProps;

export const CheckboxGroup: React.FC<CheckboxProps> = ({
    setFieldValue,
    id,
    value,
    options,
    ...rest
}) => {
    const onChange = useOnCheckboxChange(setFieldValue, id, value);
    return (
        <div className={styles.checkboxGroup}>
            {options.map((option) => {
                return (
                    <label key={option.value} className={styles.field}>
                        <input
                            key={option.value}
                            {...rest}
                            value={option.value}
                            title={option.title}
                            onChange={onChange}
                        />
                        {option.title}
                    </label>
                );
            })}
        </div>
    );
};
