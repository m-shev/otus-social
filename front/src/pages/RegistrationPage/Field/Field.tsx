import React from 'react';
import styles from './Field.module.scss';
import {
    BaseFieldProps,
    CheckboxFieldProps,
    FieldProps,
    FieldType,
    NumberFieldProps,
    SelectFieldProps,
} from './types';

const isSelectedFieldProps = (props: FieldProps): props is SelectFieldProps => {
    return props.type === FieldType.Select;
};

const isBaseFieldProps = (props: FieldProps): props is BaseFieldProps | NumberFieldProps => {
    return !([FieldType.Select, FieldType.Checkbox] as Array<string | FieldType>).includes(
        props.type,
    );
};

const isCheckboxProps = (props: FieldProps): props is CheckboxFieldProps => {
    return props.type === FieldType.Checkbox;
};

export const Field: React.FC<FieldProps> = (props) => {
    return (
        <div className={styles.root}>
            <label htmlFor={props.id} className={styles.label}>
                {props.title}
            </label>

            {isSelectedFieldProps(props) && (
                <select
                    id={props.id}
                    className={styles.field}
                    value={props.value}
                    onChange={props.onChange}
                >
                    {props.options.map((option) => {
                        return (
                            <option value={option.value} key={option.value}>
                                {option.title}
                            </option>
                        );
                    })}
                </select>
            )}

            {isBaseFieldProps(props) && <input className={styles.field} {...props} />}
        </div>
    );
};
