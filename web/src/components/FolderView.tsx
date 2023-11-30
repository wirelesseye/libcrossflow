import { ArrowUpFromLineIcon, FileIcon, FolderIcon } from "lucide-react";
import { FileInfo } from "../datatypes";
import { useFetchJSON } from "../utils/hooks";
import { FileList, FileListItem } from "./FileList";
import path from "../utils/path";

export interface FolderViewProps {
    filePath: string;
}

export default function FolderView({ filePath }: FolderViewProps) {
    const files = useFetchJSON<FileInfo[]>(`/api/files/${filePath}`);

    return (
        <FileList>
            {files !== null ? (
                <>
                    <FileListItem
                        name=".."
                        icon={<ArrowUpFromLineIcon size={20} />}
                        href={path.join("/files", path.parent(filePath))}
                    />
                    {files.map((file) => (
                        <FileListItem
                            key={file.name}
                            name={file.name}
                            href={path.join("/files/", filePath, file.name)}
                            icon={
                                file.type === "dir" ? (
                                    <FolderIcon size={20} />
                                ) : (
                                    <FileIcon size={20} />
                                )
                            }
                        />
                    ))}
                </>
            ) : null}
        </FileList>
    );
}
