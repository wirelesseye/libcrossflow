import { css, cx } from "@emotion/css";
import { HTMLProps } from "react";

export interface ButtonProps extends HTMLProps<HTMLButtonElement> {
    type?: "button" | "submit" | "reset";
    variant?: "normal" | "ghost";
}

export function Button(props: ButtonProps) {
    const { className, variant, ...other } = props;

    return (
        <button
            className={cx(
                styles.button,
                variant === "ghost" ? styles.ghost : undefined,
                className,
            )}
            {...other}
        />
    );
}

const styles = {
    button: css`
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: #ffffff;
        border: 1px solid rgba(0, 0, 0, 0.1);
        border-radius: 8px;
        padding: 8px 12px;
        transition: box-shadow 100ms, background-color 100ms;
        font-size: inherit;

        @media (hover: hover) and (pointer: fine) {
            :hover {
                box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
            }
        }

        :active {
            box-shadow: none;
            background-color: rgba(0, 0, 0, 0.2);
        }
    `,
    ghost: css`
        border: none;

        @media (hover: hover) and (pointer: fine) {
            :hover {
                background-color: rgba(0, 0, 0, 0.1);
                box-shadow: none;
            }
        }

        :active {
            background-color: rgba(0, 0, 0, 0.2);
        }
    `,
};
