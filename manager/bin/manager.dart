import 'dart:io';

import 'package:args/args.dart';
import 'package:manager/pwd_encryption.dart';
import 'package:path/path.dart' as path;

const String hashFile = "hash.json";

const String programName = "server_manager";
const String usageHeader = "Usage: $programName [arguments] <new password>";
const String newPasswordRequired =
    "The parameter <new password> is required.\n\n$usageHeader";
String createdNewPassword(String pwd, String dir) =>
    "Updated server password to: $pwd and stored $hashFile in $dir";

void main(List<String> args) {
  final ArgParser parser = initParser();

  final ArgResults results = parser.parse(args);
  if (results.wasParsed("help")) {
    print("$usageHeader\n\n${parser.usage}");
    exit(0);
  }

  if (results.rest.isEmpty) {
    print(newPasswordRequired);
    exit(1);
  }

  final String newPassword = results.rest.first;
  Directory outputDirectory = Directory.current;
  if (results.wasParsed("directory")) {
    outputDirectory = Directory(results["directory"]);
  }

  final File passwordFile = File(path.join(outputDirectory.path, hashFile))
    ..createSync();

  saveEncryptedToFile(newPassword, passwordFile);
  print(createdNewPassword(newPassword, outputDirectory.path));
  exit(0);
}

ArgParser initParser() {
  return ArgParser()
    ..addFlag(
      "help",
      abbr: "h",
      help: "Print this usage information.",
      negatable: false,
    )
    ..addOption(
      "directory",
      abbr: "d",
      help: "Write $hashFile to <directory>.",
      defaultsTo: ".",
    );
}
