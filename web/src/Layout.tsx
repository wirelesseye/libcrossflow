import { ReactNode } from "react";
import { css } from "@emotion/css";

export interface LayoutProps {
    children?: ReactNode;
}

export default function Layout({ children }: LayoutProps) {
    return (
        <div
            className={css`
                max-width: 1280px;
                margin: auto;
            `}
        >
            {children}
        </div>
    );
}
