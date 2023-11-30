import { useMemo } from "react";
import { usePathname } from "../Router";
import { useFetchJSON } from "../utils/hooks";
import { FileInfo } from "../datatypes";
import { FileList, FileListItem } from "../components/FileList";
import { ArrowUpFromLineIcon, FileIcon, FolderIcon } from "lucide-react";
import path from "../utils/path";

export default function FilesPage() {
    const pathname = usePathname();
    const filepath = useMemo(
        () => pathname.replace(/^\/files\//, ""),
        [pathname],
    );
    const files = useFetchJSON<FileInfo[]>(`/api/files/${filepath}`);

    return (
        <FileList>
            {files !== null ? (
                <>
                    <FileListItem
                        name=".."
                        icon={<ArrowUpFromLineIcon size={20} />}
                        href={path.join("/files", path.parent(filepath))}
                    />
                    {files.map((file) => (
                        <FileListItem
                            key={file.name}
                            name={file.name}
                            href={
                                file.type === "dir"
                                    ? path.join("/files/", filepath, file.name)
                                    : undefined
                            }
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
