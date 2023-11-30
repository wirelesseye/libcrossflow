function join(...args: string[]) {
    if (args.length <= 1) {
        return args.join();
    } else if (args.length === 2) {
        const l = args[0].endsWith("/")
            ? args[0].substring(0, args[0].length - 1)
            : args[0];
        const r = args[1].startsWith("/") ? args[1].substring(1) : args[1];
        return l + "/" + r;
    } else {
        const l = join(...args.slice(0, args.length - 1));
        const r = args[args.length - 1];
        return join(l, r);
    }
}

function parent(path: string) {
    if (path.endsWith("/")) {
        return parent(path.substring(0, path.length - 1));
    }
    return path.substring(0, path.lastIndexOf("/"));
}

const path = { join, parent };
export default path;
