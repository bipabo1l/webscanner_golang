var HtmlEncode = {
    Run: function (str) {
        var s = "";
        if (typeof(str) == "boolean") return str;
        if (typeof(str) == "object") return JSON.stringify(str);
        if (typeof(str) == "number") return str;
        if (str.length == 0) return "";
        s = str.replace(/</g, "&lt;");
        s = s.replace(/>/g, "&gt");
        s = s.replace(/  /g, "");
        s = s.replace(/\'/g, "\'");
        s = s.replace(/\"/g, "\"");
        s = s.replace(/\n/g, " <br>");
        return s;
    },
    IsJson: function (str) {
        if (typeof str == 'string') {
            try {
                var obj = JSON.parse(str);
                if (str.indexOf('{') > -1 || str.indexOf('[') > -1) {
                    return true;
                } else {
                    return false;
                }
            } catch (e) {
                return false;
            }
        }
        return false;
    }
};