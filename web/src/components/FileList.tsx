import { ReactNode } from "react";
import { css } from "@emotion/css";
import { Link } from "../Router";

export interface FileListProps {
    children?: ReactNode;
}

export function FileList({ children }: FileListProps) {
    return <div className={styles.fileList}>{children}</div>;
}

export interface FileListItemProps {
    icon?: ReactNode;
    name: string;
    href?: string;
}

export function FileListItem(props: FileListItemProps) {
    return (
        <Link href={props.href} className={styles.fileListItem}>
            <div className={styles.fileListIcon}>{props.icon}</div>
            {props.name}
        </Link>
    );
}

const styles = {
    fileList: css`
        display: flex;
        flex-direction: column;
    `,
    fileListItem: css`
        display: flex;
        align-items: center;
        padding: 15px 20px;
        border-radius: 8px;
        user-select: none;
        transition: background-color 100ms, box-shadow 100ms;

        :nth-child(even) {
            background-color: rgba(0, 0, 0, 0.05);
        }

        :hover {
            box-shadow: 0 3px 5px rgba(0, 0, 0, 0.2);
        }

        :active {
            box-shadow: none;
            background-color: rgba(0, 0, 0, 0.2);
        }
    `,
    fileListIcon: css`
        display: flex;
        margin-right: 10px;
    `,
};
