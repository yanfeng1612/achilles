

public class StackTest{

    private int count;

    private void recurse(){
        count ++;
        recurse();
    }

    public static void main(String[] args) {
        StackTest stackTest = new StackTest();
        try{
            stackTest.recurse();
        } catch(Throwable e){
            System.out.println("count : " + stackTest.count + "!");
        }
    }

}