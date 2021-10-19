import * as React from 'react';
import {SyntheticEvent} from 'react';
import {
    BaseFieldProps,
    CheckboxFieldProps,
    FieldProps,
    FieldType,
    NumberFieldProps,
    SelectFieldProps,
} from './types';
import styles from './Field.module.scss';

export type SpecificFieldProps = FieldProps;

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

export const SpecificField: React.FC<SpecificFieldProps> = (props) => {
    if (isSelectedFieldProps(props)) {
        return (
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
        );
    }

    if (isCheckboxProps(props)) {
        const onChange = (e: SyntheticEvent<HTMLInputElement>) => {
            const {value, setFieldValue} = props;
            const newValue = Array.isArray(value) ? [...value] : [];
            if (e.currentTarget.checked) {
                newValue.push(e.currentTarget.value);
                setFieldValue(newValue);
            } else {
                setFieldValue(
                    newValue.filter((item) => {
                        return item === e.currentTarget.value;
                    }),
                );
            }
        };
        return (
            <div className={styles.checkboxGroup}>
                {props.options.map((option) => {
                    return (
                        <label key={option.value} className={styles.field}>
                            <input
                                key={option.value}
                                {...props}
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
    }

    return <input className={styles.field} {...props} />;
};
