import { useMemo } from "react";
import { usePathname } from "../Router";
import { useFetchJSON } from "../utils/hooks";
import { FileInfo } from "../datatypes";
import { FileList, FileListItem } from "../components/FileList";
import { FileIcon, FolderIcon } from "lucide-react";

export default function FilesPage() {
    const pathname = usePathname();
    const filepath = useMemo(
        () => pathname.replace(/^\/files\//, ""),
        [pathname],
    );
    const files = useFetchJSON<FileInfo[]>(`/api/files/${filepath}`);

    return (
        <FileList>
            {files
                ? files.map((file) => (
                      <FileListItem
                          key={file.name}
                          name={file.name}
                          href={
                              file.type === "dir"
                                  ? `/files/${filepath}/${file.name}`
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
                  ))
                : null}
        </FileList>
    );
}
