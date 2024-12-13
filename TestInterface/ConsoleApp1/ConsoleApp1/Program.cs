using System.Diagnostics;

class Program
{
    static void Main(string[] args)
    {
        Process process = new Process();
        process.StartInfo.FileName = "C:\\Users\\Etudiant2\\AppData\\Local\\Discord\\Update.exe";
        process.Start();
        process.WaitForExit(); // Attendre la fin
    }
}
